package main

// Convert DMIPO to AMIPOIndividual
func convertAMIPOIndividualToDMIPO(aMIPOIndividual AMIPOIndividual) DMIPO {
	return DMIPO{
		StartDate:   aMIPOIndividual.StartDate,
		GmpUrl:      aMIPOIndividual.GmpUrl,
		Link:        aMIPOIndividual.Link,
		EndDate:     aMIPOIndividual.EndDate,
		LogoUrl:     aMIPOIndividual.LogoUrl,
		ListingDate: aMIPOIndividual.ListingDate,
		PriceRange:  aMIPOIndividual.PriceRange,
		Symbol:      aMIPOIndividual.Symbol,
		Name:        aMIPOIndividual.Name,
		Slug:        aMIPOIndividual.Slug,
	}
}

// Convert DMIPO and SMIPOIndividual to AMIPOIndividual
func mergeDMIPOAndSMIPOIndividualToAMIPOIndividual(dmIPO DMIPO, sMIPOIndividual SMIPOIndividual) AMIPOIndividual {
	return AMIPOIndividual{
		StartDate:   dmIPO.StartDate,
		GmpUrl:      dmIPO.GmpUrl,
		Link:        dmIPO.Link,
		EndDate:     dmIPO.EndDate,
		LogoUrl:     dmIPO.LogoUrl,
		ListingDate: dmIPO.ListingDate,
		PriceRange:  dmIPO.PriceRange,
		Symbol:      dmIPO.Symbol,
		Name:        dmIPO.Name,
		Slug:        dmIPO.Slug,
		Details:     sMIPOIndividual.Details,
		GmpTimeline: sMIPOIndividual.GmpTimeline,
	}
}
