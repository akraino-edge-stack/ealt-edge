package org.edgegallery.example_app.service;

import java.util.ArrayList;
import java.util.List;
import org.edgegallery.example_app.model.EALTEdgeBackup;
import org.springframework.beans.factory.annotation.Autowired;
import org.edgegallery.example_app.util.ShellCommand;
import org.springframework.stereotype.Service;

@Service
public class backupService {

    @Autowired
    private ShellCommand ShellCommands;

    public String create_backup(String backupname, String namespace) {
        String command = "velero backup create " + backupname + " --include-namespaces " + namespace;

        String output = ShellCommands.executeCommand(command);

        System.out.println(output);
        return "success";
    }

    public List<EALTEdgeBackup> getBackupTables() {

        EALTEdgeBackup backup = new EALTEdgeBackup();
        String command = "velero get backups";

        String output = ShellCommands.executeCommand(command);

        //System.out.println(output);
        List<EALTEdgeBackup> backupsList = new ArrayList<EALTEdgeBackup>();

        String list = ShellCommands.parseResult(output);

        //TODO: after parse the result, need to fill info in backup node in list
        backup.setName("backup1");

        backupsList.add(backup);

        return backupsList;
    }
}
