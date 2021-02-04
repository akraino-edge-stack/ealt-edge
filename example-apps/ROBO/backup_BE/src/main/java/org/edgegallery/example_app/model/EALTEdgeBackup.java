package org.edgegallery.example_app.model;


import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import lombok.ToString;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
public class EALTEdgeBackup {
    private String name;
    private String status;
    private String errors;
    private String warnings;
    private String created;
}
