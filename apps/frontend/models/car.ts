import type { User } from './user'

export type BaseCarModel = {
  Id: number
  CreatedAt: string
  UpdatedAt: string
  DeletedAt?: string | null
  Name: string
  BasePower: number
  BaseHandling: number
  CarInstances: Car[]
}

export type Car = {
  Id: number
  CreatedAt: string
  UpdatedAt: string
  DeletedAt?: string | null
  UserID: number
  BaseCarModelID: number
  Nickname: string
  User: User
  BaseCarModel: BaseCarModel
  InstalledComponents: InstalledCarComponent[]
}

export type CarComponent = {
  Id: number
  CreatedAt: string
  UpdatedAt: string
  DeletedAt?: string | null
  Name: string
  Description: string
  Type: string
  PowerModifier: number
  HandlingModifier: number
}

export type InstalledCarComponent = {
  Id: number
  CreatedAt: string
  UpdatedAt: string
  DeletedAt?: string | null
  CarID: number
  CarComponentID: number
  CarComponent: CarComponent
}

export type CarDTO = {
  baseCarModelId: number
  nickname?: string
}
