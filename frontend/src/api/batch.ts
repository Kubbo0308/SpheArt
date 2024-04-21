import { CONST } from "@/const";

export async function runQiitaBatch() {
  await fetch(`${CONST.API_BASE_PATH}/api${CONST.QIITA_BATCH}`)
}

export async function runZennBatch() {
  await fetch(`${CONST.API_BASE_PATH}/api${CONST.ZENN_BATCH}`)
}
