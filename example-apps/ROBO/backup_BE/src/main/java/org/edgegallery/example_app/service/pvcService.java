package org.edgegallery.example_app.service;

import io.kubernetes.client.ApiException;
import io.kubernetes.client.apis.CoreV1Api;
import io.kubernetes.client.models.V1PersistentVolumeClaim;
import io.kubernetes.client.models.V1PersistentVolumeClaimList;

import java.util.ArrayList;
import java.util.List;
import org.edgegallery.example_app.model.Pvcs;
import org.springframework.stereotype.Service;

@Service
public class pvcService {
    private Pvcs pvcsDetail;

    public List<Pvcs> getPvcsList() {
        String namespace = "default";
        CoreV1Api api = new CoreV1Api();
        V1PersistentVolumeClaimList list = null;
        try {
           /* list = api.listNamespacedPersistentVolumeClaim(namespace, null, null, null,
                    null, null, null, null, null,
                    null);*/
            list = api.listPersistentVolumeClaimForAllNamespaces(null, null, null,
                    null,null,null,null,null,null);
        } catch (ApiException apie) {
            System.err.println("Exception when calling CoreV1Api#listNamespacedPersistentVolumeClaim");
            apie.printStackTrace();
            System.exit(1);
        }

        if (list == null) {
            System.out.println("Inside- pvcs obj is null");
        }
        List<Pvcs> pvcslistElement = new ArrayList<Pvcs>();

        for (V1PersistentVolumeClaim item : list.getItems()) {
            pvcsDetail = new Pvcs();
            pvcsDetail.setNamespace(item.getMetadata().getNamespace());
            pvcsDetail.setName(item.getMetadata().getName());
            pvcsDetail.setStatus(item.getStatus().getPhase());

            pvcsDetail.setVolume(item.getSpec().getVolumeName());
            pvcsDetail.setVolumemode(item.getSpec().getVolumeMode());
           //TODO: getAccessModes return list of string, need to check
            // pvcsDetail.setAccessmodes(item.getSpec().getAccessModes());
            pvcsDetail.setStorageclass(item.getSpec().getStorageClassName());

          //  pvcsDetail.setAge(item.getStatus().getConditions().getLastTransitionTime());
            //TODO: getCapacity is a map. need to get quantity from map and fill
            //pvcsDetail.setCapacity(item.getStatus().getCapacity());
            pvcsDetail.setAge("null");
            pvcslistElement.add(pvcsDetail);
        }

        return pvcslistElement;
    }

}
