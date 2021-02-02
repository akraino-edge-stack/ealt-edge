package org.edgegallery.example_app.service;

import io.kubernetes.client.ApiException;
import io.kubernetes.client.apis.CoreV1Api;
import io.kubernetes.client.models.V1PodList;
import io.kubernetes.client.models.V1Status;
import org.springframework.stereotype.Service;
import io.kubernetes.client.models.V1DeleteOptions;

@Service
public class deleteNamespaceImp implements deleteNamespace {

    @Override
    public String deleteNS() throws ApiException {
        String namespace = "test";
        CoreV1Api api = new CoreV1Api();

        // invokes the CoreV1Api client
        //TODO: this API delete the namespace but it crash after execute. need to check some example for this API to
        // cnfirm the para
        api.deleteNamespace(namespace, null, new V1DeleteOptions(),null,null,null,null);

        System.out.println("k8s api is ok");

        return "success";
    }
}
