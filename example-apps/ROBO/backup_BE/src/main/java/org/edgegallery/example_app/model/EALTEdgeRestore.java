package org.edgegallery.example_app.model;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
public class EALTEdgeRestore {
    private String name;
    private String backup;
    private String status;
}
