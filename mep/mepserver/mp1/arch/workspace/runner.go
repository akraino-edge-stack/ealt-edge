/*
 * Copyright 2020 Huawei Technologies Co., Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package workspace

import (
	"sync"

	"mepserver/mp1/arch/bus"
)

type GoPolicy int

const (
	_ GoPolicy = iota
	GoParallel
	GoBackground
	GoSerial
)

func WkRun(plan SpaceIf) ErrCode {
	curPlan := plan.getPlan()
	for {
		if curPlan.CurGrpIdx >= len(curPlan.PlanGrp) {
			break
		}
		curSub := &curPlan.PlanGrp[curPlan.CurGrpIdx]
		retCode, stepIdx := GrpRun(curSub, plan, &curPlan.WtPlan)
		if retCode <= TaskOK {
			curPlan.CurGrpIdx++
			continue
		}
		RecordErrInfo(curPlan, stepIdx)
		GotoErrorProc(curPlan)
	}
	// wait backgroud job finish
	curPlan.WtPlan.Wait()
	return TaskOK

}

func TaskRunner(wkSpace interface{}, stepIf TaskBaseIf) int {
	for {
		bus.LoadObjByInd(stepIf, wkSpace, "in")
		retCode := stepIf.OnRequest("")
		if retCode <= TaskFinish {
			bus.LoadObjByInd(stepIf, wkSpace, "out")
			break
		}
	}
	return 0
}

func StepPolicy(wg *sync.WaitGroup, curSub *SubGrp, plan SpaceIf, wtPlan *sync.WaitGroup, stepIf TaskBaseIf) ErrCode {
	taskRet := TaskOK
	switch curSub.Policy {
	case GoBackground:
		wtPlan.Add(1)
		go func() {
			defer wtPlan.Done()
			TaskRunner(plan, stepIf)
		}()

	case GoParallel:
		wg.Add(1)
		go func() {
			defer wg.Done()
			TaskRunner(plan, stepIf)
		}()
	default:
		TaskRunner(plan, stepIf)
		taskRet, _ = stepIf.GetErrCode()
	}

	return taskRet
}

func GrpOneStep(wg *sync.WaitGroup, curSub *SubGrp, plan SpaceIf, wtPlan *sync.WaitGroup) bool {
	if curSub.CurStepIdx >= len(curSub.StepObjs) {
		return false
	}
	curStep := curSub.StepObjs[curSub.CurStepIdx]
	if curStep == nil {
		curSub.CurStepIdx++
		return true
	}
	stepIf, ok := curStep.(TaskBaseIf)
	if !ok {
		return false
	}
	taskRet := StepPolicy(wg, curSub, plan, wtPlan, stepIf)
	curSub.CurStepIdx++

	return taskRet <= TaskOK
}

func GrpGetRetCode(curSub *SubGrp) (ErrCode, int) {
	for idx, curStep := range curSub.StepObjs {
		stepIf, ok := curStep.(TaskBaseIf)
		if !ok {
			continue
		}
		errCode, _ := stepIf.GetErrCode()
		if errCode > TaskOK {
			return errCode, idx
		}
	}

	return TaskOK, -1
}

func GrpRun(curSub *SubGrp, plan SpaceIf, wtPlan *sync.WaitGroup) (ErrCode, int) {
	var wg sync.WaitGroup
	for {
		nextStep := GrpOneStep(&wg, curSub, plan, wtPlan)
		if !nextStep {
			break
		}
	}
	if curSub.Policy == GoParallel {
		wg.Wait()
	}
	return GrpGetRetCode(curSub)
}
