#!/bin/bash -eu
# 環境変数設定
BACKEND_CONTAINER_NAME=$1
BACKEND_ROOT=backend
MOCK_DATA_ROOT="testutils/mock"

# backend配下のモック化したいファイルのあるディレクトリ一覧
SEARCH_DIR_LIST="
domain/repository
usecase
"

echo "Start..."

cd ${BACKEND_ROOT}

for dir in ${SEARCH_DIR_LIST}; do
  # test.go以外のgoファイルを対象のディレクトリから探す
  file_path_list=$(find ./${dir} -type f -not -name "*_test.go" -name "*.go")

  for file_path in ${file_path_list}; do
    # file_pathからモック作成
    mockgen -package mock -source=./$file_path -destination=./$MOCK_DATA_ROOT/$file_path
  done
done

echo "Done."
