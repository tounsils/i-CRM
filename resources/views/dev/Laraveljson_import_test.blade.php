@extends('layouts.main')

@section('title', 'Laravel Json')

@section('content')


<section id="about" class="about d-flex align-items-center">
    
    <div class="container">
    
    <br>
    <br>
        
        <h1>Laravel Json</h1>

           
        <h2>Display :</h2>
        
        <p><pre>You've got  {{count($items)}}  records in the json file  </pre></p>



<br>


</div>
</section>


@endsection