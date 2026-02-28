import { Button, Form, Input } from 'antd';
import { LockOutlined, LoginOutlined, MailOutlined, UserOutlined } from '@ant-design/icons';
import { useState } from 'react';
import type { PageProps } from "keycloakify/login/pages/PageProps";
import type { I18n } from '../../i18n';
import type { KcContext } from '../../../kc.gen';
import imageLogin from '../../assets/loginImage.svg'
import './style.css'
type RegisterProps = PageProps<
  Extract<KcContext, { pageId: "register.ftl" }>,
  I18n
>;

export default function CustomRegister(props: RegisterProps) {
    const { kcContext, i18n, Template } = props;
    const [loading, setLoading] = useState(false);
    const { url, messagesPerField } = kcContext;

    return (
        <Template i18n={i18n} kcContext={kcContext} doUseDefaultCss={false} headerNode={false}>
            <div className="justify-center items-center h-screen flex">
                <section className="w-[90%] sm:w-[70%] flex rounded-[20px] bg-[#1F1F1F] overflow-hidden shadow-2xl">
                    <div className="w-[40%] hidden sm:flex">
                        <img
                            src={imageLogin}
                            alt="register background"
                            className="w-full h-full object-cover rounded-l-[20px]"
                        />
                    </div>
                    <div className="w-[100%] flex flex-col justify-center px-10 py-12 sm:w-[60%]">
                        <h1 className="text-2xl font-bold text-white mb-1">WELCOME!</h1>
                        <form
                            id="kc-form-register"
                            action={url.registrationAction}
                            method="post"
                            onSubmit={() => setLoading(true)}
                            className="flex flex-col gap-3"
                        >
                            <p className="text-white text-sm mb-1">
                                Username
                            </p>
                            <Form.Item
                                name="username"
                                rules={[{ required: true, message: "Please input your username" }]}
                                validateStatus={messagesPerField.existsError("username") ? "error" : ""}
                                help={messagesPerField.getFirstError("username")}
                                className="!mb-2"
                            >
                                <Input
                                    id="username"
                                    name="username"
                                    autoComplete="username"
                                    placeholder={'Username'}
                                    className="!h-12 !bg-[#2a2a2a] !border-gray-600 !text-white placeholder:!text-gray-500"
                                    prefix={<UserOutlined className="!text-gray-400 pr-1.5 text-lg" />}
                                />
                            </Form.Item>
                            <div className='flex gap-5'>
                                <div className='w-[50%]'>
                                    <p className="text-white text-sm mb-1">First Name</p>
                                    <Form.Item
                                        name="firstName"
                                        rules={[{ required: true, message: "Please input your firstName" }]}
                                        validateStatus={messagesPerField.existsError("firstName") ? "error" : ""}
                                        help={messagesPerField.getFirstError("firstName")}
                                        className="!mb-2"
                                    >
                                        <Input
                                            id="firstName"
                                            name="firstName"
                                            placeholder="First Name"
                                            className="!h-12 !bg-[#2a2a2a] !border-gray-600 !text-white placeholder:!text-gray-500"
                                            prefix={<UserOutlined className="!text-gray-400 pr-1.5 text-lg" />}
                                        />
                                    </Form.Item>
                                </div>
                                <div className='w-[50%]'>
                                    <p className="text-white text-sm mb-1">Last Name</p>
                                    <Form.Item
                                        name="lastName"
                                        rules={[{ required: true, message: "Please input your last name" }]}
                                        validateStatus={messagesPerField.existsError("lastName") ? "error" : ""}
                                        help={messagesPerField.getFirstError("lastName")}
                                        className="!mb-2"
                                    >
                                        <Input
                                            id="lastName"
                                            name="lastName"
                                            placeholder="Last name"
                                            className="!h-12 !bg-[#2a2a2a] !border-gray-600 !text-white placeholder:!text-gray-500"
                                            prefix={<UserOutlined className="!text-gray-400 pr-1.5 text-lg" />}
                                        />
                                    </Form.Item>
                                </div>
                            </div>
                            <p className="text-white text-sm mb-1">
                                Email
                            </p>
                            <Form.Item
                                name="username"
                                rules={[{ required: true, message: "Please input your email" }]}
                                validateStatus={messagesPerField.existsError("email") ? "error" : ""}
                                help={messagesPerField.getFirstError("email")}
                                className="!mb-2"
                            >
                                <Input
                                    id="email"
                                    name="email"
                                    type={'email'}
                                    placeholder={'Email'}
                                    className="!h-12 !bg-[#2a2a2a] !border-gray-600 !text-white placeholder:!text-gray-500"
                                    prefix={<MailOutlined className="!text-gray-400 pr-1.5 text-lg" />}
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
                            <p className="text-white text-sm mb-1">Repeat Password</p>
                            <Form.Item
                                name="password-confirm"
                                rules={[{ required: true, message: "Please input your password again" }]}
                                validateStatus={messagesPerField.existsError("password") ? "error" : ""}
                                help={messagesPerField.getFirstError("password")}
                                className="!mb-2"
                            >
                                <Input.Password
                                    id="password-confirm"
                                    name="password-confirm"
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

                        
                            <p className="text-center text-gray-400 text-sm mt-2">
                                {"You have an account? "}
                                <a href={url.loginUrl} className="text-blue-400 hover:text-blue-300 font-medium">
                                    Sign in
                                </a>
                            </p>
                        </form>
                    </div>

                </section>
            </div>
        </Template>
    );
}