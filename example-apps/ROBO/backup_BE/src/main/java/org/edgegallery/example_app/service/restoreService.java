package org.edgegallery.example_app.service;

import java.util.ArrayList;
import java.util.List;
import org.edgegallery.example_app.model.EALTEdgeBackup;
import org.edgegallery.example_app.model.EALTEdgeRestore;
import org.edgegallery.example_app.util.ShellCommand;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class restoreService {

    @Autowired
    ShellCommand shellCommand;

    public String create_restore(String restorename, String backupname) {
        String command = "velero restore create " + restorename + " --from-backup " + backupname;

        String output = shellCommand.executeCommand(command);

        System.out.println(output);
        return "success";
    }

    public List<EALTEdgeRestore> getRestoreTables() {
        EALTEdgeRestore restoreDetails = new EALTEdgeRestore();
        String command = "velero get restores";

        String output = shellCommand.executeCommand(command);

        //System.out.println(output);
        List<EALTEdgeRestore> restoresList = new ArrayList<EALTEdgeRestore>();

        String list = shellCommand.parseResult(output);

        //TODO: after parse the result, need to fill info in backup node in list
        restoreDetails.setName("restore1");

        restoresList.add(restoreDetails);
        return restoresList;
    }
}
