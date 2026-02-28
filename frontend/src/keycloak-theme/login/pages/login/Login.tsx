import { Button, Form, Input } from 'antd';
import { LockOutlined, LoginOutlined, UserOutlined } from '@ant-design/icons';
import { useState } from 'react';
import type { PageProps } from "keycloakify/login/pages/PageProps";
import type { I18n } from '../../i18n';
import type { KcContext } from '../../../kc.gen';
import imageLogin from '../../assets/loginImage.svg'
import './style.css'
type LoginProps = PageProps<
  Extract<KcContext, { pageId: "login.ftl" }>,
  I18n
>;

export default function CustomLogin(props: LoginProps) {
    const { kcContext, i18n, Template } = props;
    const [loading, setLoading] = useState(false);
    const { url, realm, messagesPerField } = kcContext;
    const { msgStr } = i18n;
    const { loginWithEmailAllowed } = realm;

    return (
        <Template i18n={i18n} kcContext={kcContext} doUseDefaultCss={false} headerNode={false}>
            <div className="justify-center items-center h-screen flex">
                <section className="w-[90%] sm:w-[70%] sm:h-[600px] h-[500px] flex rounded-[20px] bg-[#1F1F1F] overflow-hidden shadow-2xl">
                    <div className="w-[40%] hidden sm:flex">
                        <img
                            src={imageLogin}
                            alt="Login background"
                            className="w-full h-full object-cover rounded-l-[20px]"
                        />
                    </div>
                    <div className="w-[100%] flex flex-col justify-center px-10 py-12 sm:w-[60%]">
                        <h1 className="text-2xl font-bold text-white mb-1">WELCOME!</h1>
                        <form
                            id="kc-form-login"
                            action={url.loginAction}
                            method="post"
                            onSubmit={() => setLoading(true)}
                            className="flex flex-col gap-3"
                        >
                            <p className="text-white text-sm mb-1">
                                {loginWithEmailAllowed ? "Email Address" : "Username"}
                            </p>
                            <Form.Item
                                name="username"
                                rules={[{ required: true, message: "Please input your username or email" }]}
                                validateStatus={messagesPerField.existsError("username") ? "error" : ""}
                                help={messagesPerField.getFirstError("username")}
                                className="!mb-2"
                            >
                                <Input
                                    id="username"
                                    name="username"
                                    type={loginWithEmailAllowed ? "email" : "text"}
                                    autoComplete="username"
                                    placeholder={loginWithEmailAllowed ? msgStr("usernameOrEmail") : msgStr("username")}
                                    className="!h-12 !bg-[#2a2a2a] !border-gray-600 !text-white placeholder:!text-gray-500"
                                    prefix={<UserOutlined className="!text-gray-400 pr-1.5 text-lg" />}
                                />
                            </Form.Item>

                            <p className="text-white text-sm mb-1">Password</p>
                            <Form.Item
                                name="password"
                                rules={[{ required: true, message: "Please input your password" }]}
                                validateStatus={messagesPerField.existsError("password") ? "error" : ""}
                                help={messagesPerField.getFirstError("password")}
                                className="!mb-2"
                            >
                                <Input.Password
                                    id="password"
                                    name="password"
                                    autoComplete="current-password"
                                    placeholder="••••••••"
                                    className="!h-12 !bg-[#2a2a2a] !border-gray-600 !text-white placeholder:!text-gray-500"
                                    prefix={<LockOutlined className="!text-gray-400 pr-1.5 text-lg" />}
                                />
                            </Form.Item>
                            <Button
                                htmlType="submit"
                                loading={loading}
                                type="primary"
                                icon={<LoginOutlined />}
                                style={{ background: "linear-gradient(128deg, rgba(55,22,71,1) 4%, rgba(0,102,197,1) 50%, rgba(55,22,71,1) 95%)" }}
                                className="!h-12 !border-none font-semibold"
                            >
                                Sign In
                            </Button>

                        
                            {realm.registrationAllowed && (
                                <p className="text-center text-gray-400 text-sm mt-2">
                                    {"Don't have an account? "}
                                    <a href={url.registrationUrl} className="text-blue-400 hover:text-blue-300 font-medium">
                                        Sign up now
                                    </a>
                                </p>
                            )}
                        </form>
                    </div>

                </section>
            </div>
        </Template>
    );
}