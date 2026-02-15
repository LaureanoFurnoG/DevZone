import { notification } from "antd"
import type { IconType } from "antd/es/notification/interface"

export const useAppNotification = () => {
  const [api, contextHolder] = notification.useNotification()

  const notify = (title: string, description: string, type: IconType | undefined) => {
    api.open({
      title: title,
      description,
      showProgress: true,
      pauseOnHover: false,
      placement: 'bottomRight',
      type: type,
    })
  }

  return { notify, contextHolder }
}
