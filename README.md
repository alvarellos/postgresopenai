# postgresopenai


## Prerequisites

If not installed already, please ensure the latest version of [Go](https://go.dev/doc/install) is installed:

```shell
-- Install go
https://confluence.bit.admin.ch/display/itsol/Mac+setup#Macsetup-InstallGO

go install bitbucket.bit.admin.ch/en5-framework/gg@latest
```



## Installation

- Install Python

```shell
brew install python
python --version
```

- Install podman and podman-compose
```shell
python3 -m pip install --upgrade pip
python3 -m pip install podman-compose
```

- Verify installation
```shell
podman-compose --help
```

- The configuration of the podman-compose is included in the file. Feel free to change the parameters
```shell
podman_compose.yaml
```

- In order to persist the data you can create the volumes like
```shell
podman volume create data
podman volume create export
```

- The podman machine needs to be started if it has not been done already
```shell
podman machine init
```

- Start the container 
```shell
podman-compose -f ./podman_compose.yaml up
```


## Edit Global variables
- Add a key in your profile
```shell
https://platform.openai.com/api-keys
```

- The key needs to be in the global variables (use the editor that you want)
```shell
nano ~/.zshrc
```
- Include OpenAI key
```shell
export OPENAI_API_KEY=your_key
```

- Include LocalDB string
```shell
export LOCALDB=postgres://postgres:postgres@localhost:5432/postgres
```

- Optional: Include variables if you want to access the DB with terminal
```shell
export POSTGRES_USERNAME=postgres
export POSTGRES_PASSWORD=postgres
export HOSTNAME=localhost
export PORT=5432
export DATABASENAME=postgres
```

- Optional. Access with psql
```shell
psql postgres://${POSTGRES_USERNAME}:${POSTGRES_PASSWORD}@${HOSTNAME}:${PORT}/${DATABASENAME}
```
- or if the variables are not set
```shell
psql postgres://postgres:postgres@localhost:5432/postgres
```

- You are ready to build and run the program

## After running the program

- Check values
```shell
SELECT * FROM public.foo1 LIMIT 1000;
SELECT * FROM information_schema.columns WHERE table_schema = 'public' AND table_name = 'foo1';
```