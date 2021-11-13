
$festivals = App\Models\User::with(['years' => function($query) {
    return $query->select()->selectRaw("MONTHNAME(STR_TO_DATE(month_number, '%m')) as month_name");
}, 'years.city.country'])->Current()->get();