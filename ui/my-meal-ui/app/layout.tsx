import type { Metadata } from "next";
import { Inter } from "next/font/google";
import styles from "./layout.module.css";
import { AppRouterCacheProvider } from "@mui/material-nextjs/v13-appRouter";
import { ThemeProvider } from "@mui/material/styles";
import theme from "./theme";

//const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "MyMeal",
  description: "MyMeal is you choicing the best meal for you.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={styles.outer}>
        <AppRouterCacheProvider>
          <ThemeProvider theme={theme}>{children}</ThemeProvider>
        </AppRouterCacheProvider>
      </body>
    </html>
  );
}
