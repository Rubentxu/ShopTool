environments {
    local {
        repositoryCentral = "http://172.26.215.154:31006/repository/industrializacion-maven-group/"

        server {
            host = '-'
            port = 0
            fileRsa = '-'
            user = '-'
        }

    }

    dev {
        repositoryCentral = "http://rm.sva.itbatera.ejgv.eus:8081/repository/industrializacion-maven-group/"

        server {
            host = '-'
            port = 0
            fileRsa = '-'
            user = '-'
        }



    }

    prod {
        repositoryCentral = "http://rm.sva.itbatera.ejgv.eus:8081/repository/industrializacion-maven-group/"

        server {
            host = 'v99srh-000277.mbkudea.ejgv.eus'
            port = 22
            fileRsa = 'vrorun_id_rsa'
            user = 'vrorun'
        }


    }
}
