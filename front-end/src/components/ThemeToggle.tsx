import { Moon, Sun } from "lucide-react";
import { Button } from "@/components/ui/button";
import { useTheme } from "@/hooks/use-theme";

interface ThemeToggleProps {
  iconSize: string;
}

export function ThemeToggle({ iconSize }: ThemeToggleProps) {
  const { theme, setTheme } = useTheme();

  function handleToggle() {
    const newTheme = theme === "light" ? "dark" : "light";

    setTheme(newTheme);
  }

  const sizeName = `size-${iconSize}`;

  return (
    <Button
      variant="ghost"
      size="icon"
      onClick={handleToggle}
      className={`rounded-full ${sizeName}`}
      aria-label="Toggle theme"
      asChild
    >
      {theme === "light" ? (
        <Moon className={sizeName} />
      ) : (
        <Sun className={sizeName} />
      )}
    </Button>
  );
}
