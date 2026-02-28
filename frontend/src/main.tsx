import { createRoot } from "react-dom/client";
import { StrictMode, lazy, Suspense } from "react";
import { KcPage, type KcContext } from "./keycloak-theme/kc.gen";
import './index.css'

const Application = lazy(() => import("./main.app"));

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    {window.kcContext ? (
      <KcPage kcContext={window.kcContext} />
    ) : (
      <Suspense>
        <Application />
      </Suspense>
    )}
  </StrictMode>
);

declare global {
  interface Window {
    kcContext?: KcContext;
  }
}