export type User = {
  Id: number
  CreatedAt: string
  UpdatedAt: string
  DeletedAt?: string | null
  Username: string
  Email: string
  Provider: string
  ProviderID: string
  AvatarURL: string
  ProfileBio: string
  Score: number
  RacesWon: number
  TotalRaces: number
  Cars: any[]
}
