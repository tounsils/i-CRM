@extends('layouts.main')

@section('title', 'Laravel Json')

@section('content')


<section id="about" class="about d-flex align-items-center">
    
    <div class="container">
    
    <br>
    <br>
        
        <h1>Laravel Json</h1>

           
        <h2>Steps :</h2>
        
        <ul>
        <li>Choose json file</li>

        <p><pre>Laravel Json</pre></p>
        
        <li>Read from storage : challenge.json (Import it to remote folder)</li>
                <p>jsonimportController
                    <pre>
                        $path = storage_path()."/challenge.json";
                        $items_json = file_get_content($path);
                        $items = json_decode($items_json, true);
                        return view('Laraveljson_import_test', compact('items'));
                    
                    </pre></p>

        <li>Display it</li>
        <li>Choose Table</li>
        <li>Verify filelds</li>
        <li>Import to table</li>
        <li>Confirmation</li>
</ul>     

<p><pre>Laravel Json</pre></p>


<br>


</div>
</section>


@endsection