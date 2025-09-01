set allow-duplicate-variables := true

import? '.just/local.just'
import '.just/default.just'
import '.just/mod/ci.just'

[working-directory("internal/cmd/ascend-toolkit")]
tidy:
    go mod tidy

fork:
    rm -rf ./target/mind-cluster
    git clone -b v7.1.RC1 --depth=1 https://gitcode.com/Ascend/mind-cluster.git target/mind-cluster
    cp -rf ./target/mind-cluster/component/ascend-common ./internal/ascend-common
