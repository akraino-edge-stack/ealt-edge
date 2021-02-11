package org.edgegallery.example_app.service;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;
import java.util.StringTokenizer;

import org.apache.commons.lang.StringUtils;
import org.edgegallery.example_app.model.EALTEdgeBackup;
import org.springframework.beans.factory.annotation.Autowired;
import org.edgegallery.example_app.util.ShellCommand;
import org.springframework.stereotype.Service;
import org.edgegallery.example_app.common.*;

@Service
public class backupService {

    @Autowired
    private ShellCommand ShellCommands;

    public String create_backup(String backupname, String namespace) {
        String ip = System.getenv("HOSTIP");
        String command =
                "sshpass ssh root@" + ip + " velero backup create " + backupname + " --include-namespaces " + namespace;

        String output = ShellCommands.executeCommand(command);

        System.out.println(output);
        return "success";
    }

    public List<EALTEdgeBackup> getBackupTables() {
        String ip = System.getenv("HOSTIP");
        String command = "sshpass ssh root@" + ip + " velero get backups";

        List<EALTEdgeBackup> backupsList = new ArrayList<EALTEdgeBackup>();
        backupsList = ShellCommands.executeBackupCommand(command);

        return backupsList;
    }
}
