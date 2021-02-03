package org.edgegallery.example_app.model;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
public class Pvcs {
    private String namespace;
    private String name;
    private String status;
    private String volume;
    private String capacity;
    private String accessmodes;
    private String storageclass;
    private String age;
    private String volumemode;
}
