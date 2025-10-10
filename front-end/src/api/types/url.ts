export interface ShortenRequest {
  label: string;
  url: string;
}

export interface ShortenResponse {
  id: string;
}

export interface RedirectResponse {
  url: string;
}
