# NgGoKuber

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 14.2.0.

## Install

## Steps
1. Node
```shell
brew update  
brew install node
node -v
```
2. Angular 14.2.0
```shell
npm uninstall -g @angular/cli
npm cache clean
npm cache verify
sudo npm install -g @angular/cli@14.2.0
ng --version
```
3. [Tailwindcss](https://tailwindcss.com/docs/installation)
```shell
npm install -D tailwindcss
npx tailwindcss init 
```
+ configure css and path
4. Create project
```shell
ng new ng_go_kuber
```
5. [eslint](https://eslint.org/docs/latest/use/getting-started)
```shell
npm install --save-dev eslint
touch .eslintrc.js
npx eslint
eslint ./tests/** 
```
6. git
```shell
git init
git branch -m master main
git remote add origin git@github.com:vvuri/ng_go_kuber.git
git push -u origin main
```
7. makefile
8. golang api
9. docker file build
10. docker compose
11. kind
12. Prometheus
13. Grafana
14. Add metrics
15. Auth JWT
16. API test TS -> Mocha  
```shell
npm install --save-dev mocha
 npm install --save-dev ts-node --force
 npm install --save-dev @types/mocha --force
```


## Development server

Run `ng serve` for a dev server. Navigate to `http://localhost:4200/`. The application will automatically reload if you change any of the source files.

## Code scaffolding

Run `ng generate component component-name` to generate a new component. You can also use `ng generate directive|pipe|service|class|guard|interface|enum|module`.

## Build
- Run `npx eslint`
- Run `ng build` to build the project. 
- The build artifacts will be stored in the `dist/` directory.

## Running unit tests

Run `ng test` to execute the unit tests via [Karma](https://karma-runner.github.io).

## Running end-to-end tests

Run `ng e2e` to execute the end-to-end tests via a platform of your choice. To use this command, you need to first add a package that implements end-to-end testing capabilities.

## Further help

To get more help on the Angular CLI use `ng help` or go check out the [Angular CLI Overview and Command Reference](https://angular.io/cli) page.
