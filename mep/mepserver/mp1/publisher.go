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

package mp1

import (
	"sync"
	"time"

	"github.com/apache/servicecomb-service-center/pkg/gopool"
	"golang.org/x/net/context"
)

type Publisher struct {
	wss       []*Websocket
	goroutine *gopool.Pool
	lock      sync.Mutex
}

func (p *Publisher) Run() {
	gopool.Go(publisher.loop)
}

func (p *Publisher) loop(ctx context.Context) {
	defer p.Stop()
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			var removes []int
			for i, ws := range p.wss {
				payload := ws.Pick()
				if payload == nil {
					continue
				}
				_, ok := payload.(error)
				if ok {
					removes = append(removes, i)
				}
				p.dispatch(ws, payload)
			}
			if len(removes) == 0 {
				continue
			}
			p.lock.Lock()
			var (
				news []*Websocket
				s    int
			)
			for _, e := range removes {
				news = append(news, p.wss[s:e]...)
				s = e + 1
			}
			if s < len(p.wss) {
				news = append(news, p.wss[s:]...)
			}
			p.wss = news
			p.lock.Unlock()
		}
	}
}

func (p *Publisher) Stop() {
	p.goroutine.Close(true)
}

func (p *Publisher) dispatch(ws *Websocket, payload interface{}) {
	p.goroutine.Do(func(ctx context.Context) {
		ws.HandleWatchWebSocketJob(payload)
	})
}

func (p *Publisher) Accept(ws *Websocket) {
	p.lock.Lock()
	p.wss = append(p.wss, ws)
	p.lock.Unlock()
}

var publisher *Publisher

func init() {
	publisher = NewPublisher()
	publisher.Run()
}

func NewPublisher() *Publisher {
	return &Publisher{
		goroutine: gopool.New(context.Background()),
	}
}
