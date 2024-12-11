import { ReactElement, ReactNode } from "react";

interface IRoute {
  path: string;
  element: ReactElement;
  children?: IRoute[];
}
