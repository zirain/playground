#!/bin/sh


WITH_PROMETHEUS=${WITH_PROMETHEUS:-'false'}


if [ $WITH_PROMETHEUS == 'true' ]; then
    echo "$WITH_PROMETHEUS"
fi