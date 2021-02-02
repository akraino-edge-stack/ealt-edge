package org.edgegallery.example_app.service;

import io.kubernetes.client.ApiException;

public interface deleteNamespace {

    String deleteNS() throws ApiException;
}
