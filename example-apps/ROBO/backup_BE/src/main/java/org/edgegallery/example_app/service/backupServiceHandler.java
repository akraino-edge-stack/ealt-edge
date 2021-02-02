package org.edgegallery.example_app.service;

import java.util.ArrayList;
import java.util.List;
import org.edgegallery.example_app.model.EALTEdgeBackup;
import org.edgegallery.example_app.model.EALTEdgeBackupRestore;
import org.edgegallery.example_app.model.EALTEdgeRestore;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;

@Service
public class backupServiceHandler {

    @Autowired
    private backupService BackupService;

    @Autowired
    private restoreService RestoreService;
    /**
     * get back/restore tables.
     * @return
     */
    public ResponseEntity<EALTEdgeBackupRestore> getBackupRestoreDetails() {

        EALTEdgeBackupRestore eALTEdgeBackupRestore = new EALTEdgeBackupRestore();

        List<EALTEdgeBackup> backupsList = BackupService.getBackupTables();
        List<EALTEdgeRestore> restoresList = RestoreService.getRestoreTables();

        eALTEdgeBackupRestore.setBackupsData(backupsList);
        eALTEdgeBackupRestore.setRestoresData(restoresList);
        return ResponseEntity.ok(eALTEdgeBackupRestore);
    }

    /**
     * create restore tables.
     * @param restoreName restore name.
     * @param backupName backup name.
     * @return
     */
    public String createRestore(String restoreName, String backupName){
        return RestoreService.create_restore(restoreName, backupName);
    }

    /**
     * create backup tables.
     * @param backupName restore name.
     * @param namespaces backup name.
     * @return
     */
    public String createBackup(String backupName, String namespaces){
        return BackupService.create_backup(backupName, namespaces);
    }
}
