export default {
  api: {
    operationFailed: "작업 실패",
    errorTip: "에러 팁",
    errorMessage: "작업이 실패했습니다. 시스템이 정상적이지 않습니다!",
    timeoutMessage: "로그인 시간이 초과되었습니다. 다시 로그인해 주세요.",
    apiTimeoutMessage: "인터페이스 요청 시간이 초과되었습니다. 페이지를 새로고침하고 다시 시도해주세요!",
    apiRequestFailed: "인터페이스 요청이 실패했습니다. 나중에 다시 시도해 주세요!",
    networkException: "네트워크 이상",
    networkExceptionMsg: "네트워크 연결이 정상인지 확인해주세요! 네트워크가 정상적이지 않습니다.",
    errMsg401: "사용자에게 권한이 없습니다 (토큰, 사용자 이름, 비밀번호 오류)!",
    errMsg403: "사용자는 인증되었으나 접근이 금지되어 있습니다!",
    errMsg404: "네트워크 요청 오류, 리소스를 찾을 수 없습니다!",
    errMsg405: "네트워크 요청 오류, 요청 방법이 허용되지 않습니다!",
    errMsg408: "네트워크 요청 시간이 초과되었습니다!",
    errMsg500: "서버 오류, 관리자에게 문의하십시오!",
    errMsg501: "네트워크가 실행되지 않았습니다!",
    errMsg502: "네트워크 오류!",
    errMsg503: "서비스 이용이 불가능합니다. 서버가 일시적으로 과부하 상태이거나 유지보수 중입니다!",
    errMsg504: "네트워크 시간 초과!",
    errMsg505: "http 버전은 요청을 지원하지 않습니다!"
  },
  app: {
    logoutTip: "알림",
    logoutMessage: "시스템을 종료하시겠습니까?",
    menuLoading: "메뉴 로딩 중..."
  },
  errorLog: {
    tableTitle: "에러 로그 목록",
    tableColumnType: "유형",
    tableColumnDate: "시간",
    tableColumnFile: "파일",
    tableColumnMsg: "오류 메세지",
    tableColumnStackMsg: "스택 정보",
    tableActionDesc: "세부사항",
    modalTitle: "오류 세부사항",
    fireVueError: "Fire vue 에러",
    fireResourceError: "Fire 리소스 에러",
    fireAjaxError: "Fire ajax 에러",
    enableMessage: "`/src/settings/projectSetting.ts`에서 useErrorHandle=true인 경우에만 유효합니다."
  },
  exception: {
    backLogin: "로그인 돌아가기",
    backHome: "홈 돌아가기",
    subTitle403: "죄송합니다. 이 페이지에 액세스할 수 없습니다.",
    subTitle404: "죄송합니다. 방문한 페이지가 존재하지 않습니다.",
    subTitle500: "죄송합니다. 서버에서 오류를 보고하고 있습니다.",
    noDataTitle: "현재 페이지에 데이터가 없습니다.",
    networkErrorTitle: "네트워크 오류",
    networkErrorSubTitle: "죄송합니다. 네트워크 연결이 끊어졌습니다. 네트워크를 확인해주세요!"
  },
  lock: {
    unlock: "클릭하여 잠금 해제",
    alert: "잠금 화면 비밀번호 오류",
    backToLogin: "로그인으로 돌아가기",
    entry: "시스템 입장",
    placeholder: "잠금 화면 비밀번호 또는 사용자 비밀번호를 입력하세요."
  },
  login: {
    backSignIn: "로그인 돌아가기",
    mobileSignInFormTitle: "모바일 \b로그인",
    qrSignInFormTitle: "QR코드 sign in",
    signInFormTitle: 'Sign in',
    signUpFormTitle: "가입하기",
    forgetFormTitle: "비밀번호 재설정",
    signInTitle: 'SophLiteOS',
    signInDesc: "개인 정보를 입력하고 시작하세요!",
    policy: "xxx 개인정보 보호정책에 동의합니다.",
    scanSign: `코드를 스캔하여 로그인을 완료하세요`,
    loginButton: "로그인",
    registerButton: "가입하기",
    rememberMe: "기억하기",
    forgetPassword: "비밀번호를 잊으셨나요?",
    otherSignIn: "다른 방식으로 로그인",
    // notify
    loginSuccessTitle: "로그인 성공",
    loginSuccessDesc: "돌아온 것을 환영합니다.",
    // placeholder
    accountPlaceholder: "사용자 이름을 입력하세요.",
    passwordPlaceholder: "비밀번호를 입력하세요.",
    smsPlaceholder: "SMS 코드를 입력하세요.",
    mobilePlaceholder: "핸드폰 번호를 입력하세요.",
    policyPlaceholder: "확인 후 등록",
    diffPwd: "두 비밀번호가 일치하지 않습니다.",
    userName: "사용자명",
    password: "비밀번호",
    confirmPassword: "비밀번호 확인",
    email: "이메일",
    smsCode: "SMS 코드",
    mobile: "모바일"
  },
  btn: {
    confirm: "확인",
    clear: "지우기",
    reset: "리셋"
  },
  form: {
    placeholder: "입력해주세요",
    phNumber: "0~100 사이의 정수를 입력하세요."
  },
  table: {
    selectAll: "모두 확인",
    selected: "선택",
    records: "레코드(cross-page)",
    unSelect: "선택된 항목이 없습니다",
    getTableData: "테이블 데이터 가져오기",
    startTime: "시작 시간: ",
    endTime: "종료 시간: "
  },
  boardStatus: {
    error: "실패",
    running: "실행 중"
  },
  uploadFile: {
    btnText: "업로드 파일: ",
    update: "업그레이드",
    upload: "업로드",
    upPackage: "업그레이드 패키지"
  },
  copySuccess: "복사 성공",
  tips: "경고"
};