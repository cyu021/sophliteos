export default {
  api: {
    operationFailed: "操作を完了できませんでした。",
    errorTip: "エラーヒント",
    errorMessage: "操作に失敗しました。システムが異常です。",
    timeoutMessage: "ログインがタイムアウトしました。もう一度ログインしてください。",
    apiTimeoutMessage: "インターフェース リクエストがタイムアウトしました。ページを更新してもう一度お試しください",
    apiRequestFailed: "インターフェースのリクエストに失敗しました。後でもう一度お試しください",
    networkException: "ネットワーク異常",
    networkExceptionMsg: "ネットワーク接続が正常かどうかを確認してください。ネットワークが異常です",
    errMsg401: "ユーザーに権限がありません(トークン、ユーザー名、パスワードエラー)",
    errMsg403: "ユーザーは承認されましたが、アクセスは禁止されています",
    errMsg404: "ネットワークリクエストエラー。リソースが見つかりませんでした。",
    errMsg405: "ネットワークリクエストエラー。リクエストメソッドが許可されていません。",
    errMsg408: "ネットワークリクエストはタイムアウトしました",
    errMsg500: "サーバーエラー。管理者に連絡してください。",
    errMsg501: "ネットワークは実装されていません。",
    errMsg502: "ネットワークエラー",
    errMsg503: "サービスが利用できません。サーバーは一時的に過負荷またはメンテナンス中です",
    errMsg504: "ネットワークタイムアウト",
    errMsg505: "httpバージョンはリクエストをサポートしていません"
  },
  app: {
    logoutTip: "リマインダー",
    logoutMessage: "システムを終了しますか？",
    menuLoading: "メニューを読み込み中..."
  },
  errorLog: {
    tableTitle: "エラーログ一覧",
    tableColumnType: "タイプ",
    tableColumnDate: "時間",
    tableColumnFile: "ファイル",
    tableColumnMsg: "エラーメッセージ",
    tableColumnStackMsg: "スタック情報",
    tableActionDesc: "詳細",
    modalTitle: "エラーの詳細",
    fireVueError: 'Fire vue error',
    fireResourceError: "ファイヤーリソースエラー",
    fireAjaxError: 'Fire ajax error',
    enableMessage: "`/src/settings/projectSettings.ts` で useErrorHandle=true の場合にのみ有効です。"
  },
  exception: {
    backLogin: "ログインに戻る",
    backHome: "ホームへ戻る",
    subTitle403: "申し訳ありません、このページへのアクセス権がありません",
    subTitle404: "申し訳ありませんが、ご訪問いただいたページは存在しません。",
    subTitle500: "申し訳ありませんが、サーバーがエラーを報告しています。",
    noDataTitle: "現在のページにデータがありません。",
    networkErrorTitle: "ネットワークエラー",
    networkErrorSubTitle: "申し訳ありませんが、ネットワーク接続が切断されました。ネットワークを確認してください"
  },
  lock: {
    unlock: "クリックしてロック解除",
    alert: "ロック解除パスワードが無効です",
    backToLogin: "ログイン画面に戻る",
    entry: "システムを入る",
    placeholder: "ロック画面のパスワードまたはユーザーパスワードを入力してください"
  },
  login: {
    backSignIn: "サインインに戻る",
    mobileSignInFormTitle: "モバイルログイン",
    qrSignInFormTitle: "QRコードサインイン",
    signInFormTitle: "ログイン",
    signUpFormTitle: "登録",
    forgetFormTitle: "パスワードのリセット",
    signInTitle: "SophLiteOS",
    signInDesc: "個人情報を入力して始めましょう",
    policy: "[プライバシーポリシー] に同意します",
    scanSign: `コードをスキャンしてログインを完了する`,
    loginButton: "ログイン",
    registerButton: "登録",
    rememberMe: "ログイン情報を記憶する",
    forgetPassword: "パスワードを忘れましたか?",
    otherSignIn: "次でログイン",
    // notify
    loginSuccessTitle: "ログインしました",
    loginSuccessDesc: "おかえりなさい",
    // placeholder
    accountPlaceholder: "ユーザー名を入力下さい",
    passwordPlaceholder: "パスワードを入力してください",
    smsPlaceholder: "SMSコードを入力してください",
    mobilePlaceholder: "携帯電話を入力してください",
    policyPlaceholder: "確認後に登録",
    diffPwd: "パスワードが一致していません",
    userName: "ユーザー名",
    password: "パスワード",
    confirmPassword: "パスワード再確認",
    email: "メール",
    smsCode: "SMSコード",
    mobile: "携帯電話"
  },
  btn: {
    confirm: "確認",
    clear: "クリア",
    reset: "リセット"
  },
  form: {
    placeholder: "入力してください。",
    phNumber: "0から100までの整数を入力してください"
  },
  table: {
    selectAll: "すべてを確認",
    selected: "選択済",
    records: "レコード(クロスページ)",
    unSelect: "アイテムが選択されていません",
    getTableData: "テーブルデータを取得",
    startTime: "開始時間",
    endTime: "終了時間"
  },
  boardStatus: {
    error: "失敗",
    running: "実行中"
  },
  uploadFile: {
    btnText: "ファイルアップロード: ",
    update: "アップグレード",
    upload: "アップロード",
    upPackage: "アップグレードパッケージ"
  },
  copySuccess: "コピー完了",
  tips: "警告"
};