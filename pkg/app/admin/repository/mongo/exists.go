package mongo

func (r *Repository) Exists(ctx context.Context, roleAdminID string) (bool, error) {

	oID, err := bson.ObjectIDFromHex(roleAdminID)
	if err != nil {
		return false, err
	}

	filter := bson.M{
		"roles_ids": bson.M{
			"$in": []bson.ObjectID{oID},
		},
	}

	count, err := r.UserColl.CountDocuments(ctx, filter)
	if err != nil {
	  return  false, err
	}
	return count > 0, nil

}