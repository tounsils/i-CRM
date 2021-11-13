@extends('layouts.main')

@section('title', 'Laravel Database seeding')

@section('content')


<section id="about" class="about d-flex align-items-center">
    
    <div class="container">



        <h1>Tinker</h1>

        
<p><pre>php artisan tinker</pre></p>
<p><pre>App\Models\User::find(1)</pre></p>
<p><pre>
=> App\Models\user {#4349
     id: 1,
     first_name: "Ms. Elody Littel",
     last_name: "Reynolds",
     email: "tofgfggfs@gmail.com",
     phone: "(848) 769-9406",
     role_id: 1,
     #password: "$2y$10$jWpmrx.aL.4VztHsSxPLuuJTxUmAkrd0CaVXSPpCwWYRCLZkhB4Bm",
     #remember_token: null,
     created_at: "2021-07-24 21:43:48",
     updated_at: "2021-07-24 21:43:48",
   }    
</pre></p>
<p><pre>
>>> App\Models\User::find(2)->role
=> App\Models\role {#4352
     id: 2,
     name: "Author",
     slug: "author",
     created_at: null,
     updated_at: null,
   }

</pre></p>

<p><pre>>>> App\Models\role::find(1)->users</pre></p>
 Dispalys all users having role = 2


<p><pre>
=> Illuminate\Database\Eloquent\Collection {#4392 
     all: [
       App\Models\user {#4393
         id: 2,
         first_name: "Dorian Douglas",
         last_name: "Raynor",
         email: "kevin.kulas@example.net",
         phone: "+1-779-258-5609",
         role_id: 2,
         #password: "{\,,Y,_",
         #remember_token: null,
         created_at: "2021-07-24 21:43:49",
         updated_at: "2021-07-24 21:43:49",
       },
       App\Models\user {#4394
         id: 3,
         first_name: "Wellington Hessel",
         last_name: "Fay",
         email: "herman.aliyah@example.com",
         phone: "1-570-652-1700",
         role_id: 2,
         #password: ",bpvmjAMB`7T~ycw",
         #remember_token: null,
         created_at: "2021-07-24 21:43:49",
         updated_at: "2021-07-24 21:43:49",
       },
       App\Models\user {#4395
         id: 4,
         first_name: "Paige Rice I",
         last_name: "Price",
         email: "jammie11@example.com",
         phone: "341.635.9700",
         role_id: 2,
         #password: "d=9\jn5dZ_M1J",
         #remember_token: null,
         created_at: "2021-07-24 21:43:49",
         updated_at: "2021-07-24 21:43:49",
       },
       App\Models\user {#4396
         id: 5,
         first_name: "Clarissa Leffler",
         last_name: "Schaefer",
         email: "crona.elliott@example.net",
         phone: "+1 (272) 950-5883",
         role_id: 2,
         #password: "uxGKfE6z6,JG",
         #remember_token: null,
         created_at: "2021-07-25 13:46:00",
         updated_at: "2021-07-25 13:46:00",
       },
       App\Models\user {#4397
         id: 6,
         first_name: "Godfrey McKenzie",
         last_name: "Crist",
         email: "ronaldo.kuhn@example.net",
         phone: "1-959-863-8635",
         role_id: 2,
         #password: "0T4SP7]W$",
         #remember_token: null,
         created_at: "2021-07-25 13:46:01",
         updated_at: "2021-07-25 13:46:01",
       },
       App\Models\user {#4398
         id: 7,
         first_name: "Kaylin Labadie",
         last_name: "Schuster",
         email: "vanessa.buckridge@example.org",
         phone: "+18209270835",
         role_id: 2,
         #password: "{u59PWY", .......    

</pre></p>



<h2>Introduction</h2>
<p>
    
</p>

<br>


</div>
</section>


@endsection