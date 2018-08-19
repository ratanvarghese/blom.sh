for dir in ./*; do
	[ -e "$dir/item.json" ] || continue
	sh bin/article.sh -A "$dir" -T ./template
done

items=$(cat */item.json \
		| jq -s 'sort_by(.date_published) | reverse | del(.[]._ratan_blog_class)')

jq -n -c -f template/jfroot.jq.json \
	--argjson ITEMS "$items" \
	| tee feeds/json \
	| gzip --keep --best > feeds/json.gz

