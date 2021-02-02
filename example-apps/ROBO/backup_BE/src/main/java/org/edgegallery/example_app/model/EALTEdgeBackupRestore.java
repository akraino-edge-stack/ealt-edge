package org.edgegallery.example_app.model;

import java.util.List;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import lombok.ToString;
import org.springframework.stereotype.Service;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
@Service
public class EALTEdgeBackupRestore {

    private List<EALTEdgeBackup> backupsData;
    private List<EALTEdgeRestore> restoresData;

}
