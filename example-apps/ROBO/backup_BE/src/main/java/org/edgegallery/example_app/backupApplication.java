package org.edgegallery.example_app;


import io.kubernetes.client.ApiException;
import io.kubernetes.client.Configuration;
import io.kubernetes.client.apis.CoreV1Api;
import io.kubernetes.client.models.V1PersistentVolumeClaim;
import io.kubernetes.client.models.V1PersistentVolumeClaimList;
import io.kubernetes.client.util.ClientBuilder;
import io.kubernetes.client.util.KubeConfig;
import java.util.ArrayList;
import java.util.List;
import org.edgegallery.example_app.model.Pvcs;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import io.kubernetes.client.ApiClient;
import org.springframework.context.annotation.Bean;
import java.io.FileReader;
import java.io.IOException;

@SpringBootApplication
public class backupApplication {

    @Bean
    public static void apiclient() throws IOException {
        // file path to your KubeConfig
        //String homePath = System.getenv("HOME");
        //String kubeConfigPath = homePath + "/.kube/config";
        String kubeConfigPath = System.getenv("KUBE_CONFIG");

        // loading the out-of-cluster config, a kubeconfig from file-system
        ApiClient client =
                ClientBuilder.kubeconfig(KubeConfig.loadKubeConfig(new FileReader(kubeConfigPath))).build();

        // set the global default api-client to the in-cluster one from above
        Configuration.setDefaultApiClient(client);
    }

    public static void main(String[] args) throws IOException {
        SpringApplication.run(backupApplication.class, args);
    }
}