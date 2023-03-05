import webapi from "./gocliRequest"
import * as components from "./fizzbuzzComponents"
export * from "./fizzbuzzComponents"

const basrUrl = "http://localhost:8888"

/**
 * @description 
 * @param req
 */
export function fizzbuzz(req: components.FizzbuzzRequest) {
	return webapi.post<components.FizzbuzzResponse>(`${basrUrl}/fizzbuzz`, req)
}
