import "./Login.css"

const Login = () => {
    const onClickGoogle = event => {
        event.preventDefault();
        window.location.assign('http://localhost:8000/auth/google/login');
    }

    const onClickGithub = event => {
        event.preventDefault();
        window.location.assign('http://localhost:8000/auth/github/login');
    }
    return (
        <div>
        <button className="google-button" onClick={onClickGoogle}>Google</button>
        <button className="github-button" onClick={onClickGithub}>Github</button>
        </div>
    )
}

export default Login