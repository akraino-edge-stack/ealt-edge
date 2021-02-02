package org.edgegallery.example_app.controller;

import io.swagger.annotations.ApiOperation;
import io.swagger.annotations.ApiResponse;
import io.swagger.annotations.ApiResponses;
import javax.ws.rs.core.MediaType;
import org.edgegallery.example_app.model.*;
import io.kubernetes.client.ApiException;
import io.kubernetes.client.apis.CoreV1Api;
import io.kubernetes.client.models.*;
import org.checkerframework.common.reflection.qual.GetMethod;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.*;
import org.edgegallery.example_app.service.podService;
import org.edgegallery.example_app.service.pvcService;
import org.edgegallery.example_app.service.deleteNamespace;


import java.util.ArrayList;
import java.util.List;

@CrossOrigin
@Controller
@RequestMapping("/v1/robo")
@Validated
public class k8sController {

    @Autowired
    EALTEdgePodsPvcs eALTEdgePodsPvcs;

    @Autowired
    pvcService PvcService;

    @Autowired
    podService PodService;

    private final deleteNamespace DeleteNamespace;

    @Autowired
    public k8sController(deleteNamespace DeleteNamespace) {
        this.DeleteNamespace = DeleteNamespace;
    }

    @GetMapping(path = "/apps-pvcs")
    @ApiOperation(value = "get pod and pvcs tables.", response = EALTEdgePodsPvcs.class,
            responseContainer = "List")
    @ApiResponses(value = {
            @ApiResponse(code = 404, message = "microservice not found", response = String.class),
            @ApiResponse(code = 415, message = "Unprocessable " + "MicroServiceInfo Entity ",
                    response = String.class),
            @ApiResponse(code = 500, message = "resource grant " + "error", response = String.class)
    })
    public ResponseEntity<EALTEdgePodsPvcs> getAllPodPvcList() throws ApiException {

        List<Pod> podlistElement;
        podlistElement = PodService.getPodsList();
        if (podlistElement.isEmpty()) {
            System.out.println("Pod list is null");
        }

        List<Pvcs> pvcslistElement;
        pvcslistElement = PvcService.getPvcsList();
        if (pvcslistElement.isEmpty()) {
            System.out.println("Pvcs list is null");
        }

        eALTEdgePodsPvcs.setAppsData(podlistElement);
        eALTEdgePodsPvcs.setPvcData(pvcslistElement);

        return new ResponseEntity<EALTEdgePodsPvcs>(eALTEdgePodsPvcs, HttpStatus.OK);
    }

    @GetMapping(path = "/disaster")
    @ApiOperation(value = "delete namespace. in k8s", response = String.class,
            responseContainer = "List")
    @ApiResponses(value = {
            @ApiResponse(code = 404, message = "microservice not found", response = String.class),
            @ApiResponse(code = 415, message = "Unprocessable " + "MicroServiceInfo Entity ",
                    response = String.class),
            @ApiResponse(code = 500, message = "resource grant " + "error", response = String.class)
    })
    public ResponseEntity<String> createDisaster() throws ApiException {

        DeleteNamespace.deleteNS();
        System.out.println("k8s 1 api is ok");

        return ResponseEntity.ok("success");
    }
}
