
const nodemailer = require("nodemailer");
// 메일발송 객체
const mailSender = {
  // 메일발송 함수
  sendGmail: function () {
    var transporter = nodemailer.createTransport({
      service: 'gmail',   // 메일 보내는 곳
      prot: 587,
      host: 'smtp.gmlail.com',  
      secure: false,  
      requireTLS: true ,
      auth: {
        user: "cujuserver@gmail.com",  // 보내는 메일의 주소
        pass: "cujuroot1!"   // 보내는 메일의 비밀번호
      }
    });
    // 메일 옵션
    var mailOptions = {
      from: "cujuserver@gmail.com", // 보내는 메일의 주소
      to: "chs29359685@mgmail.com", // 수신할 이메일
      subject: "test", // 메일 제목
      text: "Fuck" // 메일 내용
    };
    
    // 메일 발송    
    transporter.sendMail(mailOptions, function (error, info) {
      if (error) {
        console.log(error);
      } else {
        console.log('Email sent: ' + info.response);
      }
    });

  }
}

module.exports = mailSender;