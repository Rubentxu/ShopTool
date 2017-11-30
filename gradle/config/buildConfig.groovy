environments {
    local {
        repositoryCentral = "http://172.26.215.154:31006/repository/industrializacion-maven-group/"

        server {
            host = '52.174.123.255'
            port = 22
            fileRsa = 'terraform_rsa'
            user = 'rubentxu'
            passPhrase = 'cabrera'
        }

    }

    dev {
        repositoryCentral = "http://rm.sva.itbatera.ejgv.eus:8081/repository/industrializacion-maven-group/"

        server {
            host = '-'
            port = 0
            fileRsa = '-'
            user = '-'
            passPhrase = '-'
        }



    }

    prod {
        repositoryCentral = "http://rm.sva.itbatera.ejgv.eus:8081/repository/industrializacion-maven-group/"

        server {
            host = '-'
            port = 22
            fileRsa = '-'
            user = '-'
            passPhrase = '-'
        }


    }
}
