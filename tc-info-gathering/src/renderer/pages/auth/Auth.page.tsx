import { FC, useState } from "react";
import toast from 'react-hot-toast';
import './Auth.page.css';
import { Button,  TextField } from "@mui/material";
import { useNavigate } from "react-router-dom";

const CORRECT_EMAIL = 'yaroslav11@gmail.com'
const CORRECT_PASSWORD = 'qwerty12345'

export const AuthPage: FC = () => {
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const navigate = useNavigate();

    const signIn = (email: string, password: string) => {        
        if (email !== CORRECT_EMAIL || password !== CORRECT_PASSWORD) {
            toast.error('Неверный логин или пароль');
            return
        }
    
        navigate('/monitoring');
    }

    return (
        <div>
            <div className="auth-page_credentials">
                <TextField
                    id="outlined-password-input"
                    label="Email"
                    type="email"
                    value={email}
                    onChange={(event) => setEmail(event.target.value)}
                />
                <TextField
                    id="outlined-password-input"
                    label="Пароль"
                    type="password"
                    autoComplete="current-password"
                    value={password}
                    onChange={(event) => setPassword(event.target.value)}
                />
            </div>
            <div className="auth-page_buttons">
                <Button onClick={() => signIn(email, password)} color="secondary" variant="contained">Войти</Button>
                <Button color="secondary" variant="contained">Зарегистрироваться</Button>
            </div>
        </div>
    )
}
