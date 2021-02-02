package org.edgegallery.example_app.service;

import javax.validation.constraints.NotEmpty;
import javax.validation.constraints.Pattern;
import javax.validation.constraints.Size;
import org.springframework.validation.annotation.Validated;

import static org.edgegallery.example_app.common.Constants.NAME_REGEX;

import javax.validation.constraints.NotEmpty;
import javax.validation.constraints.Pattern;
import javax.validation.constraints.Size;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import lombok.ToString;
import org.springframework.validation.annotation.Validated;

/**
 * Create instance input schema.
 */
@Validated
@Getter
@Setter
@ToString
@NoArgsConstructor
@AllArgsConstructor
public class createParam {

    @NotEmpty(message = "backup name is mandatory")
    @Size(max = 64)
    @Pattern(regexp = NAME_REGEX)
    private String backupName;

    @NotEmpty(message = "namespace is mandatory")
    @Size(max = 64)
    @Pattern(regexp = NAME_REGEX)
    private String namespace;

/*    public String getBackupName() {
        return backupName;
    }

    public void setBackupName(String backupName) {
        this.backupName = backupName;
    }*/
}

