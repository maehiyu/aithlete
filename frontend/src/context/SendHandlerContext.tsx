import { createContext, useContext } from "react";

export type SendHandler = (input: string) => void;

export const SendHandlerContext = createContext<{
    handler: SendHandler;
    setHandler: (fn: SendHandler) => void;
}>({ 
    handler: () => {},
    setHandler: () => {}
});

export function useSendHandler() {
    const ctx = useContext(SendHandlerContext);
    if (!ctx) throw new Error("useSendHandler must be used within a SendHandlerContext.Provider");
    return ctx;
}
