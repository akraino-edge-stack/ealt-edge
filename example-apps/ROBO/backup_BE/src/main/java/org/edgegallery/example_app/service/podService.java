package org.edgegallery.example_app.service;

import io.kubernetes.client.ApiException;
import io.kubernetes.client.apis.CoreV1Api;
import io.kubernetes.client.models.V1Pod;
import io.kubernetes.client.models.V1PodList;
import java.util.ArrayList;
import java.util.List;
import org.edgegallery.example_app.model.Pod;
import org.springframework.stereotype.Service;

@Service
public class podService {

    public List<Pod> getPodsList() throws ApiException {

        Pod podDetail;

        // the CoreV1Api loads default api-client from global configuration.
        CoreV1Api api = new CoreV1Api();

        // invokes the CoreV1Api client

        V1PodList eliotPods = api.listPodForAllNamespaces(null, null,
                null, null, null, null, null, null, null);

        List<Pod> podlistElement = new ArrayList<Pod>();

        for (V1Pod item : eliotPods.getItems()) {
            podDetail = new Pod();
            podDetail.setNamespace(item.getMetadata().getNamespace());
            podDetail.setName(item.getMetadata().getName());
            podDetail.setStatus(item.getStatus().getPhase());
            podDetail.setIp(item.getStatus().getPodIP());
            podDetail.setNode(item.getSpec().getNodeName());
            podDetail.setReadiness("null");
            podlistElement.add(podDetail);
        }

        return podlistElement;
    }
}
