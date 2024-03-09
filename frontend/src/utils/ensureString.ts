export function EnsureString(input: string | string[] | undefined): string {
  if (typeof input === 'string') {
    return input;
  } else if (Array.isArray(input)) {
    // 配列の場合、最初の要素を返すか、配列を結合して一つの文字列にする
    return input.join(', '); // または input[0] で最初の要素を取得
  } else {
    // undefinedの場合、デフォルト値を設定
    return '';
  }
}
