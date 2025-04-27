export GF_GCFG_PATH=${GF_GCFG_PATH}
export GOOSE_DRIVER=${GOOSE_DRIVER}
export GOOSE_DBSTRING=${GOOSE_DBSTRING}
export GOOSE_MIGRATION_DIR=${GOOSE_MIGRATION_DIR}
export WLINK_DATABASE_DEFAULT_LINK=${WLINK_DATABASE_DEFAULT_LINK}
export WLINK_REDIS_DEFAULT_ADDRESS=${WLINK_REDIS_DEFAULT_ADDRESS}
export WLINK_REDIS_DEFAULT_PASS=${WLINK_REDIS_DEFAULT_PASS}
export WLINK_REDIS_CACHE_ADDRESS=${WLINK_REDIS_CACHE_ADDRESS}
export WLINK_REDIS_CACHE_PASS=${WLINK_REDIS_CACHE_PASS}
if ! command -v goose &> /dev/null; then
    echo "goose未安装，正在安装..."
    curl -fsSL https://raw.githubusercontent.com/pressly/goose/master/install.sh | sh
    if command -v goose &> /dev/null; then
        echo "goose安装成功！"
    else
        echo "安装失败，请检查网络连接或手动安装。" >&2
        exit 1
    fi
fi
goose up
if pgrep -x "wlink" >/dev/null; then
    echo "终止 wlink 进程..."
    kill -9 $(pgrep -x "wlink")
fi
nohup /home/wlinkdir/wlink  > /home/wlinkdir/output.log 2>&1 &