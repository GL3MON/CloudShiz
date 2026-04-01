provider "kubernetes" {
    config_path = "~/.kube/config"
}

resource "kubernetes_deployment" "app" {
    metadata {
        name = "app"
        labels = {
            app = "app"
        }
    }
    
    spec {
        replicas = 2

        selector {
            match_labels = {
                app = "app"
            }
        }

        template {
            metadata {
                labels = {
                    app: "app"
                }
            }

            spec {
                container {
                    name = "nginx"
                    image = "nginx"

                    port {
                        container_port = 80
                    }
                }
            }
        }
    }
}

resource "kubernetes_service" "app" {
    metadata {
        name = "app"
    }

    spec {
        selector  = {
            app = "app"
        }
        
        port {
            port = 80
            target_port = 80
        }

        type = "NodePort"
    }
}