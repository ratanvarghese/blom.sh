for dir in ./*; do
	[ -e "$dir/item.json" ] || continue
	sh bin/article.sh -A "$dir" -T ./template
done

cat */item.json | jq -s 'sort_by(.date_published) | reverse'
