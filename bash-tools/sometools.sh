#!/bin/bash

grep -lr OLD_VERSION | xargs sed -i 's#OLD_VERSION#NEW_VERSION#'
