type HTTP_STATUS_MESSAGE = {[key:number]:string}
const DEFAULT_HTTP_STATUS_MESSAGES : HTTP_STATUS_MESSAGE = {

    400: 'Bad Requests', // client에서 request 요청 시 누락 또는 잘못된 데이터 전송 및 요청
    401: 'Unauthorized',
    403: 'Forbidden',
    404: 'Not Found',
    405: 'Wrong method', // GET인데 POST로 요청하거나 .. 
    409: 'duplicate',
    500: 'Internal Server Error', // Server 로직 수행 중 터짐
    503: 'Temporary Unavailable',
};

//interface 이용해 Error 객체에 statusCode key 추가
export interface ErrorWithStatusCode extends Error {
    statusCode? : number
};

const errorGenerator = ({ msg='', statusCode=500}: { msg?: string, statusCode: number }): void => {
    //인자로 들어오는 메세지와 상태 코드를 매핑
    const err: ErrorWithStatusCode = new Error(msg || DEFAULT_HTTP_STATUS_MESSAGES[statusCode]);
    err.statusCode = statusCode;
    throw err;
}

export default errorGenerator;

