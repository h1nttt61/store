package service

import "fmt"

func GenerateEmailHTML(code string) string {
	return fmt.Sprintf(`
    <!DOCTYPE html>
    <html>
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Код подтверждения</title>
        <style>
            body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: #f4f4f9; margin: 0; padding: 0; }
            .container { max-width: 500px; margin: 40px auto; background-color: #ffffff; border-radius: 12px; box-shadow: 0 4px 15px rgba(0,0,0,0.1); overflow: hidden; }
            .header { background: linear-gradient(135deg, #6b46c1, #d53f8c); padding: 30px; text-align: center; color: white; }
            .header h1 { margin: 0; font-size: 24px; }
            .content { padding: 40px 30px; text-align: center; color: #333; }
            .content p { font-size: 16px; line-height: 1.6; color: #555; }
            .code-box { background-color: #f3e8ff; border: 2px dashed #6b46c1; border-radius: 8px; padding: 20px; margin: 30px 0; }
            .code { font-size: 36px; font-weight: bold; letter-spacing: 8px; color: #6b46c1; margin: 0; }
            .footer { background-color: #f9f9f9; padding: 20px; text-align: center; font-size: 12px; color: #999; }
            .button { display: inline-block; background-color: #6b46c1; color: white; text-decoration: none; padding: 12px 30px; border-radius: 6px; font-weight: bold; margin-top: 20px; }
        </style>
    </head>
    <body>
        <div class="container">
            <div class="header">
                <h1>🐱 Kittens Shop</h1>
            </div>
            <div class="content">
                <h2>Подтверждение входа</h2>
                <p>Привет! Мы получили запрос на вход в твой аккаунт. Вот твой код подтверждения:</p>
                <div class="code-box">
                    <p class="code">%s</p>
                </div>
                <p>Код действителен в течение 10 минут. Если ты не запрашивал вход, просто проигнорируй это письмо.</p>
            </div>
            <div class="footer">
                <p>© 2026 Kittens Shop. Все права защищены.</p>
            </div>
        </div>
    </body>
    </html>
    `, code)
}
