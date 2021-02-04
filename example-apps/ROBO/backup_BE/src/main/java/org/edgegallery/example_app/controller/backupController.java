package org.edgegallery.example_app.controller;

import io.swagger.annotations.ApiOperation;
import io.swagger.annotations.ApiParam;
import io.swagger.annotations.ApiResponse;
import io.swagger.annotations.ApiResponses;
import java.util.List;
import javax.validation.Valid;
import javax.validation.constraints.NotNull;
import javax.ws.rs.core.MediaType;
import org.hibernate.validator.constraints.Length;
import org.edgegallery.example_app.model.EALTEdgeBackupRestore;
import org.edgegallery.example_app.service.backupServiceHandler;
import org.edgegallery.example_app.service.createParam;
import org.edgegallery.example_app.service.createParamRestore;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestPart;
import org.springframework.web.bind.annotation.RestController;

@CrossOrigin
@Controller
@RequestMapping("/v1/robo")
@Validated
public class backupController {
    private static final int MAX_COMMON_STRING_LENGTH = 255;

    private static final int MAX_DETAILS_STRING_LENGTH = 1024;

    @Autowired
    private backupServiceHandler BackupServiceHandler;

    @GetMapping(value = "/backup-restore", produces = MediaType.APPLICATION_JSON)
    @ApiOperation(value = "get backup and restore tables.", response = EALTEdgeBackupRestore.class,
            responseContainer = "List")
    @ApiResponses(value = {
            @ApiResponse(code = 404, message = "microservice not found", response = String.class),
            @ApiResponse(code = 415, message = "Unprocessable " + "MicroServiceInfo Entity ",
                    response = String.class),
            @ApiResponse(code = 500, message = "resource grant " + "error", response = String.class)
    })
    public ResponseEntity<EALTEdgeBackupRestore> getBackupRestoreDetails() {
        return BackupServiceHandler.getBackupRestoreDetails();
    }

    @PostMapping(value = "/backup", produces = MediaType.APPLICATION_JSON)
    @ApiOperation(value = "create backup.", response = String.class)
    @ApiResponses(value = {
            @ApiResponse(code = 404, message = "microservice not found", response = String.class),
            @ApiResponse(code = 415, message = "Unprocessable " + "MicroServiceInfo Entity ",
                    response = String.class),
            @ApiResponse(code = 500, message = "resource grant " + "error", response = String.class)
    })
    public ResponseEntity<String> getBackupRestoreDetails(@ApiParam(value = "create backup instance")
                                                               @Valid @RequestBody createParam CreateParam) {
        BackupServiceHandler.createBackup(CreateParam.getBackupName(), CreateParam.getNamespace());
        return ResponseEntity.ok("create backup success.");
    }


/*    @PostMapping(value = "/v1/robo/backup", produces = MediaType.APPLICATION_JSON)
    public ResponseEntity<String> getBackupRestoreDetails( @ApiParam(value = "create backup instance")
                                                               @Valid @RequestBody createParam CreateParam) {
        BackupServiceHandler.createBackup(CreateParam.getBackupName(), CreateParam.getNamespace());
        return ResponseEntity.ok("create backup success.");
    }*/

    @PostMapping(value = "/restore", produces = MediaType.APPLICATION_JSON)
    @ApiOperation(value = "create restore.", response = String.class)
    @ApiResponses(value = {
            @ApiResponse(code = 404, message = "microservice not found", response = String.class),
            @ApiResponse(code = 415, message = "Unprocessable " + "MicroServiceInfo Entity ",
                    response = String.class),
            @ApiResponse(code = 500, message = "resource grant " + "error", response = String.class)
    })
    public ResponseEntity<String> createRestore(@ApiParam(value = "create restore instance")
                                                    @Valid @RequestBody createParamRestore CreateParam) {
        BackupServiceHandler.createRestore(CreateParam.getRestoreName(), CreateParam.getBackupName());
        return ResponseEntity.ok("create restore success.");
    }
}
