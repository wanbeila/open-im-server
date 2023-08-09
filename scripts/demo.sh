#!/usr/bin/env bash

clear
. $(dirname ${BASH_SOURCE})/lib/util.sh

openim::util::desc "You can learn a lot about automation using make help"
openim::util::run "make help"
clear

openim::util::desc "You can learn a lot about automation using make help-all"
openim::util::run "make help-all"
clear

openim::util::desc "Run tidy"
openim::util::run "make tidy"
clear

openim::util::desc "Vendor go.mod"
openim::util::run "make vendor"
clear

openim::util::desc "Code style: fmt, vet, lint"
openim::util::run "make style"
clear

openim::util::desc "Run unit tests"
openim::util::run "make test"
clear

openim::util::desc "Run unit tests and get test coverage"
openim::util::run "make cover"
clear

openim::util::desc "Check for updates to go.mod dependencies"
openim::util::run "make updates"
clear

openim::util::desc "Clean"
openim::util::run "make clean"
clear

openim::util::desc "Generate all necessary files"
openim::util::run "make gen"
clear

openim::util::desc "Generate swagger document"
openim::util::run "make swagger"
clear

openim::util::desc "Serve swagger spec and docs"
openim::util::run "make serve-swagger"
clear

openim::util::desc "Verify the license headers for all files"
openim::util::run "make verify-copyright"
clear

openim::util::desc "Add copyright"
openim::util::run "make add-copyright"
clear

openim::util::desc "Project introduction, become a contributor"
openim::util::run "make advertise"
clear

openim::util::desc "Release the project"
openim::util::run "make release"
clear

openim::util::desc "Run demo"
openim::util::run "make demo"
clear
