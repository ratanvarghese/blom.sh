title=''
tag_list='[]'
article_class=''
template_path=$PWD
article_path=$PWD

while getopts ":t:c:g:T:A:" opt; do
	case $opt in
		t)
			title=$OPTARG
			;;
		c)
			article_class="$OPTARG $article_class"
			;;
		g)
			tag_list=$(jq -n	--argjson LIST "$tag_list" \
								--arg TAG "$OPTARG" \
								'$LIST + [$TAG]')
			;;
		T)
			template_path=$(realpath $OPTARG)
			;;
		A)
			article_path=$(realpath $OPTARG)
			;;
		:)
			echo "Option -$OPTARG requires argument"
			exit 1
			;;
	esac
done

editday_raw=$(stat $article_path/content.md -c %z)
article_url="http://www.ratan.blog/${article_path##*/}"
today_tq=$(date | tqdate --short | tr -d "\n" | tr " " -)
today_gr=$(date --iso-8601 | tr -d "\n")
editday_item=$(date -d "$editday_raw" --rfc-3339=seconds | tr -d "\n" | tr " " T)
editday_tq=$(date -d "$editday_raw" | tqdate --short | tr -d "\n" | tr " " -)
editday_gr=$(date -d "$editday_raw" --iso-8601 | tr -d "\n")

content=$(markdown $article_path/content.md | tee $article_path/content.html)
gzip -k --best $article_path/content.md
gzip -k --best $article_path/content.html

if [ -f "$article_path/item.json" ]; then
	writeday_item=$(cat $article_path/item.json | jq -r '.date_published')
	if [ -z "$title" ]; then
		title=$(cat $article_path/item.json | jq -r '.title')
	fi
	if [ -z "$article_class" ]; then
		old_class=$(cat $article_path/item.json | jq -r '._ratan_blog_class')
		if [ "$old_class" != "null" ]; then
			article_class=$old_class
		fi
	fi
	if [ "$tag_list" = '[]' ]; then
		tag_list=$(cat $article_path/item.json | jq '.tags')
	fi
else
	writeday_item=$editday_item
fi

writeday_tq=$(date -d "$writeday_item" | tqdate --short | tr -d "\n" | tr " " -)
writeday_gr=$(date -d "$writeday_item" --iso-8601 | tr -d "\n")

attach_list='[]'
for attach in $article_path/attachments/*.*; do
	[ -e "$attach" ] || continue
	mime=$(file --mime-type -b $attach)
	url=${attach/$article_path/$article_url}
	attach_list=$(jq -n --argjson LIST "$attach_list" \
						--arg URL "$url" \
						--arg MIME "$mime" \
						'$LIST + [{url: $URL, mime_type: $MIME}]')
done

jq	-n -c -f $template_path/jfitem.jq.json \
	--arg URL $article_url \
	--arg TITLE "$title" \
	--arg EDITDAY "$editday_item" \
	--arg WRITEDAY "$writeday_item" \
	--arg CONTENT "$content" \
	--arg ARTICLECLASS "$article_class" \
	--argjson ATTACHLIST "$attach_list" \
	--argjson TAGLIST "$tag_list" > $article_path/item.json

m4	-D_TITLE="$title" \
	-D_ARTICLE_CLASS="$article_class" \
	-D_TODAY_TQ="$today_tq" \
	-D_TODAY_GR="$today_gr" \
	-D_EDITDAY_TQ="$editday_tq" \
	-D_EDITDAY_GR="$editday_gr" \
	-D_WRITEDAY_TQ="$writeday_tq" \
	-D_WRITEDAY_GR="$writeday_gr" \
	-D_ARTICLE_CONTENT="$content" \
	-D_PERMALINK="$article_url" \
	-D_MDLINK="$article_url/content.md" \
	-D_BASICLINK="$article_url/content.html" \
	$template_path/content.m4 $template_path/page.html \
	| tee $article_path/index.html \
	| gzip --best > $article_path/index.html.gz
