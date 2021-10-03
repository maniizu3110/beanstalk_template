#! /bin/sh

# non-zeroの時exitするための設定
set -e

echo "run db migration"
source /app/app.env
/app/migrate -path /app/db/migration -database "$DB_SOURCE" -verbose up

# スクリプトに渡されたすべてのパラメータを取得して実行する
echo "start the app"
exec "$@"