# Jumpstart Templates

Welcome to the Jumpstart Templates repository! This repository contains a collection of templates for various frameworks and languages to help you kickstart your projects with ease using the Jumpstart  CLI & VSCode extension.

## Contribution Guidelines

We welcome contributions to the Jumpstart Templates repository! Follow these guidelines to contribute:

1. Fork the repository to your GitHub account.
2. Create a new branch for your feature or bug fix with the following branch name convention ```feat/framework-name``` or ```fix/framework-name```.
3. Make your changes to the templates or add new templates. Ensure to update the corresponding YAML files with accurate information.
4. Commit your changes with clear and concise messages.
5. Push your changes to your forked repository.
6. Open a pull request with details about your changes. Provide a clear description of the purpose and impact of your contribution.
7. Your contribution will be reviewed, and upon approval, it will be merged into the main repository.

## Template Structure

Each template is a folder with the following structure:

```sh
template-name
├── package.yml
└── templates
    ├── README.md
    └── ...
```

### `package.yml` File Structure

The `package.yml` file contains the metadata for the template. It is a YAML file with the following structure:

```yml
name: Template name

# This is used to check if it exists on the system
binary: <binary name>

commands:
- <command 1>
- <command 2>

# Default pipeline configuration
pipeline:
  setup:
  -
  -
  build:
  -
  install:
  - 
  - 
  run:

```


Let's take a look at each of the fields in the `package.yml` file:

- `name`: The name of the template.
- `binary`: The name of the binary used to check if the template exists on the system.
- `commands`: The list of commands to run so that a project gets generated using this template.
- `pipeline`: The default pipeline configuration for the template to generate the `Dockerfile`. The pipeline is a list of stages, and each stage is a list of commands in that order.
   1. `setup`: The setup stage is used to install any dependencies required by the template.
   2. `build`: The build stage is used to build the project.
   3. `install`: The install stage is used to install any dependencies required by the project.
   4. `run`: The run stage is used to run the project.

   The `setup` stage is added before the `build` stage, and the `install` stage is added before the `run` stage.

   <table>
  <thead>
    <tr>
      <th>Pipeline</th>
      <th>Dockerfile</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><pre><code>pipeline:
  setup:
  - FROM golang
  - WORKDIR /src
  - COPY go.* ./
  - RUN go mod download
  - COPY . .
  # Perform the build
  build:
  - RUN make releases/{{name}}
  install:
  - FROM alpine
  - COPY --from=0 /src/releases/{{name}} /
  run: /{{name}}</code></pre></td>
      <td><pre><code>FROM golang
WORKDIR /src
COPY go.* ./
RUN go mod download
COPY . .
RUN make releases/gogin
FROM alpine
COPY --from=0 /src/releases/gogin /
ENTRYPOINT /gogin</code></pre></td>
    </tr>
  </tbody>
</table>

#### Template Variables

The `package.yml` template file can contain variables that are replaced with the corresponding values when the project is generated. The variables are enclosed in double curly braces `{{}}`. The following variables are available:

- `{{name}}`: The name of the project that will be generated using this template.
- `{{codeprovider}}`: The code provider is the VCS host. For now, only GitHub is supported. The default value is `github.com`.
- `{{namespace}}`: The namespace is the organization or user name on the code provider.

### License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/nurdsoft/jumpstart-templates/blob/main/LICENCE.md) file for details.

### Support

If you encounter any issues or have questions, feel free to [open an issue](https://github.com/nurdsoft/jumpstart-templates/issues). We appreciate your feedback!

Happy coding with Jumpstart Templates!